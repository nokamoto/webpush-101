package nokamoto.webpush101

import java.util.Base64
import java.util.concurrent.TimeUnit

import com.typesafe.config.ConfigFactory
import io.grpc.netty.NettyServerBuilder
import nokamoto.protobuf.WebpushServiceGrpc
import nokamoto.webpush101.lib.{KeyPair, WebpushHttpClient}

import scala.concurrent.ExecutionContext
import scala.util.Try

object Main {
  def main(args: Array[String]): Unit = {
    val config = ConfigFactory.load()
    val port = Try(config.getInt("port")).getOrElse(9000)
    val priv = Try(config.getString("priv")).getOrElse("")
    val pub = Try(config.getString("pub")).getOrElse("")
    val ctx = ExecutionContext.global
    val client = WebpushHttpClient(
      KeyPair(priv = Base64.getDecoder.decode(priv),
              pub = Base64.getDecoder.decode(pub)))
    val server = NettyServerBuilder
      .forPort(port)
      .addService(
        WebpushServiceGrpc.bindService(new WebpushService(client)(ctx), ctx))
      .build()

    println(s"*** start gRPC server $port")
    server.start()

    sys.addShutdownHook {
      println("*** shutting down gRPC server since JVM is shutting down")
      server.shutdown()
      if (!server.awaitTermination(10, TimeUnit.SECONDS)) {
        server.shutdownNow()
      }
      println("*** server shut down")
    }

    server.awaitTermination()
  }
}
