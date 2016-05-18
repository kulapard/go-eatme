class GoEatme < Formula
  desc ""
  homepage ""
  url "https://github.com/kulapard/go-eatme/archive/0.1.2.tar.gz"
  version "0.1.2"
  sha256 "efcf18fa186e4037725d5d610aff10c755201588ea62e223a4b632b78f18c379"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    (buildpath/"src/github.com/kulapard/go-eatme/").install Dir["*"]
    system "go", "build", "-o", "#{bin}/eatme", "-v", "github.com/kulapard/go-eatme/"
  end

  test do
    system "false"
  end
end
