package controllers

import io.grpc.StatusRuntimeException
import javax.inject.Inject
import javax.inject.Singleton
import models.PushSubscription
import nokamoto.protobuf.PushSubscriptionServiceGrpc.PushSubscriptionServiceStub
import nokamoto.protobuf.WebpushServiceGrpc.WebpushServiceStub
import nokamoto.protobuf.PushSubscriptionNotification
import nokamoto.protobuf.User
import nokamoto.protobuf.UserSubscription
import play.api.Logger
import play.api.libs.json.Json
import play.api.libs.json.Reads
import play.api.mvc._

import scala.concurrent.ExecutionContext
import scala.concurrent.Future
import scala.util.control.NonFatal

@Singleton
class WebpushApplication @Inject()(
    webpushServiceStub: WebpushServiceStub,
    pushSubscriptionServiceStub: PushSubscriptionServiceStub,
    cc: ControllerComponents)
    extends AbstractController(cc) {

  private[this] implicit val ec: ExecutionContext = cc.executionContext

  def index: Action[AnyContent] =
    Action(Ok(views.html.WebpushApplication.index()))

  private[this] def action[A](f: Request[A] => Future[_])(
      implicit r: Reads[A]): Action[A] = Action.async(parse.json[A]) { req =>
    f(req).map(_ => Ok(Json.obj())).recover {
      case e: StatusRuntimeException =>
        Logger.error(e.getStatus.toString)
        InternalServerError(Json.obj("error" -> e.getMessage))

      case NonFatal(e) =>
        Logger.error(e.getMessage)
        InternalServerError(Json.obj("error" -> e.getMessage))
    }
  }

  def subscribe: Action[PushSubscription] = action[PushSubscription] { req =>
    val user = User() // TODO
    val m = UserSubscription().update(_.user := user,
                                      _.subscription := req.body.proto :: Nil)
    pushSubscriptionServiceStub.subscribe(m)
  }

  def unsubscribe: Action[PushSubscription] = action[PushSubscription] { req =>
    val m = req.body.proto
    pushSubscriptionServiceStub.unsubscribe(m)
  }

  def testSubscription: Action[PushSubscription] = action[PushSubscription] {
    req =>
      val m = PushSubscriptionNotification().update(
        _.subscription := req.body.proto :: Nil,
        _.request.content := "test subscription")
      webpushServiceStub.sendPushSubscriptionNotification(m)
  }
}
