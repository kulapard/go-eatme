class GoEatme < Formula
  desc ""
  homepage ""
  url "https://github.com/kulapard/go-eatme/archive/0.1.0.tar.gz"
  version "0.1.0"
  sha256 "9e1d32b54ff0134f61505312541c1a7e8b3b2d55278aa1847071f27cd3c9f9f5"

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
