package modules

import com.google.inject.AbstractModule
import io.grpc.ManagedChannelBuilder
import nokamoto.protobuf.WebpushServiceGrpc
import nokamoto.protobuf.WebpushServiceGrpc.WebpushServiceStub
import play.api.Configuration
import play.api.Environment

class GrpcModule(environment: Environment, configuration: Configuration)
    extends AbstractModule {
  override def configure(): Unit = {
    val host = configuration.get[String]("grpc.webpush.host")
    val port = configuration.get[Int]("grpc.webpush.port")
    val channel =
      ManagedChannelBuilder.forAddress(host, port).usePlaintext().build()
    bind(classOf[WebpushServiceStub])
      .toInstance(WebpushServiceGrpc.stub(channel))
  }
}
