require 'formula'

class SysvbannerGo < Formula
  homepage 'https://github.com/winebarrel/sysvbanner-go'
  url 'https://github.com/winebarrel/sysvbanner-go/releases/download/0.1.1/sysvbanner-go-0.1.1-darwin-amd64.tar.gz'
  sha1 'a119a21bc08f11745dfad81e38a53a1d27914006'
  version '0.1.1'
  head 'https://github.com/winebarrel/sysvbanner-go.git', :branch => 'master'

  def install
    bin.install 'banner'
  end
end
