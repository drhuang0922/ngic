class Ngic < Formula
  desc "Next Generation Image Converter - Convert JPG/PNG to WebP/AVIF"
  homepage "https://github.com/drhuang0922/ngic"
  url "https://github.com/drhuang0922/ngic/archive/v1.0.0.tar.gz"
  sha256 "" # This will need to be updated when you create a release
  license "MIT"
  head "https://github.com/drhuang0922/ngic.git", branch: "main"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/ngic"
  end

  test do
    # Test version output
    assert_match version.to_s, shell_output("#{bin}/ngic -version")
    
    # Test help output
    help_output = shell_output("#{bin}/ngic -h")
    assert_match "Next Generation Image Converter", help_output
    assert_match "Usage:", help_output
  end
end
