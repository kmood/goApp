package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["context.go"],
    importpath = "github.com/DazzlingSun/monitorService/src/basic/net/rpc/liverpc/context",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = ["//library/net/rpc/liverpc:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
