# This file was generated by GoReleaser. DO NOT EDIT.
class Om < Formula
  desc ""
  homepage ""
  version "6.0.1"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/pivotal-cf/om/releases/download/6.0.1/om-darwin-6.0.1.tar.gz"
    sha256 "5cad51dc937f192b9ca6d6e25426c6ea2de2f70d714dda0183ee05071c1320c0"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/om/releases/download/6.0.1/om-linux-6.0.1.tar.gz"
      sha256 "77f92a2feca1f9c04c0a2f8c39b5d446bd3e3381481047b7434eed930928d3f8"
    end
  end

  def install
    bin.install "om"
  end

  test do
    system "#{bin}/om --version"
  end
end
