package nokamoto.webpush101

import java.util.Base64

import com.google.protobuf.ByteString
import com.google.protobuf.empty.Empty
import com.squareup.okhttp.mockwebserver.{MockResponse, MockWebServer}
import io.grpc.{Status, StatusRuntimeException}
import nokamoto.protobuf.{PushSubscription, PushSubscriptionNotification}
import nokamoto.webpush101.lib.{KeyPair, WebpushHttpClient}
import org.scalatest.{Assertion, AsyncFlatSpec, FlatSpec}

import scala.concurrent.Future

class WebpushServiceSpec extends AsyncFlatSpec {
  private[this] def decode(s: String) = Base64.getDecoder.decode(s)

  private[this] val client = WebpushHttpClient(
    pair = KeyPair(
      priv = decode("AJFotoB4FS7IX6tbm5t0SGyISTQ6l54mMzpfYipdOD+N"),
      pub = decode(
        "BNuvjW90TpDawYyxhvK79QVyNEplaSQZOWo1CwXDmWwfya6qnyBvIx3tFvKEBetExvil4rNNRL0/ZR2WLjGEAbQ=")
    ))

  private[this] val auth = decode("LsUmSxGzGt+KcuczkTfFrQ==")

  private[this] val p256dh = decode(
    "BOVFfCoBB/2Sn6YZrKytKc1asM+IOXFKz6+T1NLOnrGrRXh/xJEgiJIoFBO9I6twWDAj6OYvhval8jxq8F4K0iM=")

  private[this] def test(
      f: (MockWebServer, PushSubscriptionNotification) => Future[Assertion])
    : Future[Assertion] = {
    val server = new MockWebServer()
    server.start()

    val req = PushSubscriptionNotification().update(
      _.subscription := PushSubscription().update(
        _.endpoint := server.url("/test").toString,
        _.auth := ByteString.copyFrom(auth),
        _.p256Dh := ByteString.copyFrom(p256dh)) :: Nil)

    val result = f(server, req)
    result.onComplete(_ => server.shutdown())
    result
  }

  "WebpushService" should "return OK if push service returns 201" in {
    test { (server, req) =>
      server.enqueue(new MockResponse().setResponseCode(201))

      val sut = new WebpushService(client)

      sut.sendPushSubscriptionNotification(req).map(_ => succeed)
    }
  }

  "WebpushService" should "return UNIMPLEMENTED if push service returns 400" in {
    test { (server, req) =>
      server.enqueue(new MockResponse().setResponseCode(400))

      val sut = new WebpushService(client)

      sut
        .sendPushSubscriptionNotification(req)
        .map(_ => fail("expected error"))
        .recover {
          case e: StatusRuntimeException
              if e.getStatus == Status.UNIMPLEMENTED =>
            succeed
        }
    }
  }
}
