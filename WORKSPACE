workspace(name = "go_proto_expression")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_go",
    patch_args = ["-p1"],
    patches = [
        "//:io_bazel_rules_go.pull.3083.patch",
    ],
    remote = "https://github.com/bazelbuild/rules_go",
    tag = "v0.31.0",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:nogo",
    version = "1.18.1",
)

git_repository(
    name = "com_google_protobuf",
    branch = "main",
    remote = "https://github.com/protocolbuffers/protobuf",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

git_repository(
    name = "bazel_gazelle",
    patch_args = ["-p1"],
    patches = ["//:bazel_gazelle.issue.1217.patch"],
    remote = "https://github.com/bazelbuild/bazel-gazelle",
    tag = "v0.25.0",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_iancoleman_strcase",
    importpath = "github.com/iancoleman/strcase",
    sum = "h1:05I4QRnGpI0m37iZQRuskXh+w77mr6Z41lwQzuHLwW0=",
    version = "v0.2.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:w43yiav+6bVFTBQFZX0r7ipe9JQ1QsbMgHwbBziscLw=",
    version = "v1.28.0",
)

go_repository(
    name = "tech_einride_go_aip",
    importpath = "go.einride.tech/aip",
    sum = "h1:srys7sFWPixEqyOu0gWuZAC86p4UAnWJIQcA01Ys3R4=",
    version = "v0.54.1",
)

gazelle_dependencies()
