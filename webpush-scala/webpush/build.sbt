scalaVersion := "2.12.6"

version := IO.readLines(file("../../VERSION")).head

enablePlugins(JavaAppPackaging)

PB.protoSources in Compile := (file("../../protobuf").getCanonicalFile * AllPassFilter).get

PB.includePaths in Compile := Seq(file("../..").getCanonicalFile,
                                  target.value / "protobuf_external")

PB.targets in Compile := Seq(
  scalapb
    .gen(grpc = true, flatPackage = true) -> (sourceManaged in Compile).value
)

libraryDependencies ++= Seq(
  "com.thesamet.scalapb" %% "scalapb-runtime" % scalapb.compiler.Version.scalapbVersion % "protobuf",
  "io.grpc" % "grpc-netty" % scalapb.compiler.Version.grpcJavaVersion,
  "com.thesamet.scalapb" %% "scalapb-runtime-grpc" % scalapb.compiler.Version.scalapbVersion,
  "com.typesafe" % "config" % "1.3.3",
  "org.scalatest" %% "scalatest" % "3.0.5" % "test",
  "com.google.crypto.tink" % "apps-webpush" % "1.2.0",
  "com.auth0" % "java-jwt" % "3.4.0",
  "com.squareup.okhttp" % "okhttp" % "2.7.5",
  "com.squareup.okhttp" % "mockwebserver" % "2.7.5" % Test,
)
