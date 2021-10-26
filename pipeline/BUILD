load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

package(
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])

go_binary(
    name = "simplesum_main",
    srcs = ["simplesum_main.go"],
    deps = [
        "@com_github_apache_beam//sdks/go/pkg/beam:go_default_library",
        "@com_github_apache_beam//sdks/go/pkg/beam/io/filesystem/gcs:go_default_library",
        "@com_github_apache_beam//sdks/go/pkg/beam/io/filesystem/local:go_default_library",
        "@com_github_apache_beam//sdks/go/pkg/beam/io/textio:go_default_library",
        "@com_github_apache_beam//sdks/go/pkg/beam/log:go_default_library",
        "@com_github_apache_beam//sdks/go/pkg/beam/transforms/stats:go_default_library",
        "@com_github_apache_beam//sdks/go/pkg/beam/x/beamx:go_default_library",
    ],
)

go_binary(
    name = "command_runner",
    srcs = ["command_runner.go"],
    deps = ["@com_github_golang_glog//:go_default_library"],
)


container_image(
    name = "command_runner_image",
    base = "@base_image//image",
    creation_time = "{BUILD_TIMESTAMP}",
    entrypoint = ["/command_runner"],
    files = [
        ":command_runner",
        ":simplesum_main",
    ],
    stamp = True,
)

container_push(
    name = "command_runner_image_publish",
    format = "Docker",
    image = ":command_runner_image",
    registry = "gcr.io",
    repository = "$(GCR_PROJECT)/command_runner",
    tag = "$(VERSION)",
)
