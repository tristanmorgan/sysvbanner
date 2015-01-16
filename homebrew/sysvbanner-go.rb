require 'formula'

class SysvbannerGo < Formula
  homepage 'https://github.com/winebarrel/sysvbanner-go'
  url 'https://github.com/winebarrel/sysvbanner-go/releases/download/0.1.0/sysvbanner-go-0.1.0-darwin-amd64.tar.gz'
  sha1 '0c4ff71a441d115555b67e9076c01f006c8ad6f3'
  version '0.1.0'
  head 'https://github.com/winebarrel/sysvbanner-go.git', :branch => 'master'

  def install
    bin.install 'banner'
  end
end
