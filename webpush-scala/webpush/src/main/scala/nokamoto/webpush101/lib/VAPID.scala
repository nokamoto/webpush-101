package nokamoto.webpush101.lib

import java.net.URL
import java.time.{LocalDateTime, ZoneOffset}
import java.util.Date

import com.auth0.jwt.JWT

import scala.concurrent.duration.Duration

object VAPID {
  private[this] def t(endpoint: String,
                      sub: String,
                      expiry: Duration,
                      pair: KeyPair): String = {
    val url = new URL(endpoint)
    val aud = s"${url.getProtocol}://${url.getHost}"
    val exp = Date.from(
      LocalDateTime
        .now()
        .plusSeconds(expiry.toSeconds)
        .toInstant(ZoneOffset.UTC))
    JWT
      .create()
      .withAudience(aud)
      .withExpiresAt(exp)
      .withSubject(sub)
      .sign(pair.alg)
  }

  private[this] def k(pair: KeyPair): String = pair.k

  def vapid(endpoint: String,
            subject: String,
            expiry: Duration,
            pair: KeyPair): String =
    s"vapid t=${t(endpoint, subject, expiry, pair)},k=${k(pair)}"
}
