package controllers

import javax.inject.Inject
import javax.inject.Singleton
import models.PushSubscription
import nokamoto.protobuf.PushSubscriptionNotification
import nokamoto.protobuf.WebpushServiceGrpc.WebpushServiceStub
import play.api.libs.json.Json
import play.api.mvc.AbstractController
import play.api.mvc.Action
import play.api.mvc.AnyContent
import play.api.mvc.ControllerComponents

import scala.concurrent.ExecutionContext
import scala.util.control.NonFatal

@Singleton
class WebpushApplication @Inject()(webpushServiceStub: WebpushServiceStub,
                                   cc: ControllerComponents)
    extends AbstractController(cc) {

  private[this] implicit val ec: ExecutionContext = cc.executionContext

  def index: Action[AnyContent] =
    Action(Ok(views.html.WebpushApplication.index()))

  def subscribe: Action[AnyContent] = TODO

  def unsubscribe: Action[AnyContent] = TODO

  def testSubscription: Action[PushSubscription] =
    Action.async(parse.json[PushSubscription]) { req =>
      val m = PushSubscriptionNotification().update(
        _.subscription := req.body.proto :: Nil,
        _.request.content := "test subscription")
      webpushServiceStub
        .sendPushSubscriptionNotification(m)
        .map(_ => Ok(Json.obj()))
        .recover {
          case NonFatal(e) =>
            InternalServerError(Json.obj("error" -> e.getMessage))
        }
    }
}
