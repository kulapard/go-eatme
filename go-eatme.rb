class GoEatme < Formula
  desc ""
  homepage ""
  url "https://github.com/kulapard/go-eatme/archive/0.1.1.tar.gz"
  version "0.1.1"
  sha256 "bf56b17606921ca2cf483b87385cbca98721ee0322d32a0a621b24d0adb55252"

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
