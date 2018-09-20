package nokamoto.webpush101.lib

import java.nio.charset.Charset

import com.google.crypto.tink.apps.webpush.WebPushHybridEncrypt
import com.squareup.okhttp._
import nokamoto.protobuf.PushSubscription

import scala.concurrent.duration._

case class WebpushHttpClient(pair: KeyPair) {
  private[this] val client = new OkHttpClient()

  private[this] val utf8 = Charset.forName("UTF-8")

  private[this] val mediaType = MediaType.parse("application/octet-stream")

  def send(subscription: PushSubscription, body: String): Response = {
    val encrypt = new WebPushHybridEncrypt.Builder()
      .withAuthSecret(subscription.auth.toByteArray)
      .withRecipientPublicKey(subscription.p256Dh.toByteArray)
      .build()
    val ciphertext = encrypt.encrypt(body.getBytes(utf8),
                                     null /* contextInfo, must be null */ )
    val b = RequestBody.create(mediaType, ciphertext)

    val req = new Request.Builder()
      .url(subscription.endpoint)
      .post(b)
      .addHeader("Authorization",
                 VAPID.vapid(subscription.endpoint,
                             "mailto:nokamoto.engr@gmail.com",
                             12.hours,
                             pair))
      .addHeader("TTL", "30")
      .addHeader("Content-Encoding", "aes128gcm")
      .addHeader("Content-Length", b.contentLength().toString)
      .build()

    client.newCall(req).execute()
  }
}
