scalaVersion := "2.12.6"

enablePlugins(PlayScala)

PB.protoSources in Compile := (file("../../protobuf").getCanonicalFile * AllPassFilter).get

PB.includePaths in Compile := Seq(file("../..").getCanonicalFile,
                                  target.value / "protobuf_external")

includeFilter in PB.generate := new SimpleFileFilter(
  (f: File) => f.getPath.endsWith(".proto"))

PB.targets in Compile := Seq(
  scalapb
    .gen(grpc = true, flatPackage = true) -> (sourceManaged in Compile).value
)

libraryDependencies ++= Seq(
  guice,
  "org.scalatestplus.play" %% "scalatestplus-play" % "3.1.2" % Test,
  "org.mockito" % "mockito-core" % "1.10.19" % Test,
  "com.thesamet.scalapb" %% "scalapb-runtime" % scalapb.compiler.Version.scalapbVersion % "protobuf",
  "io.grpc" % "grpc-netty" % scalapb.compiler.Version.grpcJavaVersion,
  "com.thesamet.scalapb" %% "scalapb-runtime-grpc" % scalapb.compiler.Version.scalapbVersion
)
