package nokamoto.webpush101.lib

import java.util.Base64

import com.auth0.jwt.algorithms.Algorithm
import com.google.crypto.tink.subtle.EllipticCurves

case class KeyPair(priv: Array[Byte], pub: Array[Byte]) {
  private[lib] val k =
    Base64.getUrlEncoder.withoutPadding().encodeToString(pub)

  private[this] val key =
    EllipticCurves.getEcPrivateKey(EllipticCurves.CurveType.NIST_P256, priv)

  private[lib] val alg = Algorithm.ECDSA256(null, key)
}
