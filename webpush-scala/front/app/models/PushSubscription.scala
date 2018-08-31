package models

import java.util.Base64

import com.google.protobuf.ByteString
import play.api.libs.json.Json
import play.api.libs.json.OFormat

case class PushSubscription(endpoint: String, p256dh: String, auth: String) {
  def proto: nokamoto.protobuf.PushSubscription = {
    nokamoto.protobuf
      .PushSubscription()
      .update(_.endpoint := endpoint,
              _.p256Dh := ByteString.copyFrom(Base64.getDecoder.decode(p256dh)),
              _.auth := ByteString.copyFrom(Base64.getDecoder.decode(auth)))
  }
}

object PushSubscription {
  implicit val format: OFormat[PushSubscription] = Json.format[PushSubscription]
}
