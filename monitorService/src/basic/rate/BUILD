package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["rate.go"],
    importpath = "github.com/DazzlingSun/monitorService/src/basic/rate",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//library/rate/limit:all-srcs",
        "//library/rate/vegas:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
