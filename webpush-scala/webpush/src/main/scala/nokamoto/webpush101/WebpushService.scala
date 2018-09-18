package nokamoto.webpush101
import com.google.protobuf.empty.Empty
import com.squareup.okhttp.Response
import io.grpc.{Status, StatusRuntimeException}
import nokamoto.protobuf.PushSubscriptionNotification
import nokamoto.webpush101.lib.WebpushHttpClient

import scala.concurrent.{ExecutionContext, Future}

class WebpushService(client: WebpushHttpClient)(implicit ctx: ExecutionContext)
    extends nokamoto.protobuf.WebpushServiceGrpc.WebpushService {
  override def sendPushSubscriptionNotification(
      request: PushSubscriptionNotification): Future[Empty] = {
    val f: Future[Seq[Response]] = Future.traverse(request.subscription) {
      subscription =>
        Future.successful(client.send(subscription, request.getRequest.content))
    }

    f.flatMap { res =>
      res.find(_.code() / 100 != 2) match {
        case Some(err) =>
          println(err)
          Future.failed(new StatusRuntimeException(Status.UNIMPLEMENTED))

        case None => Future.successful(Empty())
      }
    }
  }
}
