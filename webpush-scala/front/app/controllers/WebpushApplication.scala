package controllers

import javax.inject.Inject
import play.api.mvc.AbstractController
import play.api.mvc.Action
import play.api.mvc.AnyContent
import play.api.mvc.ControllerComponents

class WebpushApplication @Inject()(cc: ControllerComponents)
    extends AbstractController(cc) {
  def index: Action[AnyContent] =
    Action(Ok(views.html.WebpushApplication.index()))

  def subscribe: Action[AnyContent] = TODO

  def unsubscribe: Action[AnyContent] = TODO

  def testSubscription: Action[AnyContent] = TODO
}
