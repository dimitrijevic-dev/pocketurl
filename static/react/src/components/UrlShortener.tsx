import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Link, ArrowRight, Copy, CheckCircle } from "lucide-react";
import { useToast } from "@/hooks/use-toast";

interface ShortenedUrl {
  id: number;
  origin_url: string;
  destination_url: string;
  expires_at: string;
  created_at: string;
}

const UrlShortener = () => {
  const [url, setUrl] = useState("");
  const [domain, setDomain] = useState("pocketurl.zip");
  const [isLoading, setIsLoading] = useState(false);
  const [shortenedUrl, setShortenedUrl] = useState<string | null>(null);
  const [copied, setCopied] = useState(false);
  const { toast } = useToast();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!url) return;
    
    setIsLoading(true);
    setShortenedUrl(null);
    
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL || 'https://api.pocketurl.zip'}/links`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          destination_url: url,
          expires_at: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString(),
          domain: domain,
        }),
      });

      if (response.ok) {
        const data: ShortenedUrl = await response.json();
        const fullShortUrl = `https://s.${domain}/${data.origin_url}`;
        setShortenedUrl(fullShortUrl);
        toast({
          title: "Success!",
          description: "Your URL has been shortened successfully.",
        });
      } else {
        throw new Error('Failed to shorten URL');
      }
    } catch (error) {
      if (process.env.NODE_ENV === 'development') {
        console.error('Error shortening URL:', error);
      }
      toast({
        title: "Error",
        description: "Failed to shorten URL. Please try again.",
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  const handleCopy = async () => {
    if (!shortenedUrl) return;
    
    try {
      await navigator.clipboard.writeText(shortenedUrl);
      setCopied(true);
      toast({
        title: "Copied!",
        description: "URL copied to clipboard",
      });
      setTimeout(() => setCopied(false), 2000);
    } catch (error) {
      if (process.env.NODE_ENV === 'development') {
        console.error('Failed to copy:', error);
      }
    }
  };

  return (
    <div className="w-full max-w-lg mx-auto">
      <Card className="backdrop-blur-xl bg-white/20 border border-white/30 shadow-2xl p-8 rounded-3xl">
        <div className="space-y-8">
          <div className="text-center space-y-3">
            <div className="inline-flex items-center justify-center w-16 h-16 bg-white/20 rounded-2xl mb-4">
              <Link className="w-8 h-8 text-gray-700" />
            </div>
            <h2 className="text-2xl font-semibold text-gray-900 tracking-tight">
              Shorten Links
            </h2>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-3">
              <label htmlFor="url" className="block text-sm font-medium text-gray-800">
                Enter your URL
              </label>
              <Input
                id="url"
                type="url"
                placeholder="https://example.com/very-long-url"
                value={url}
                onChange={(e) => setUrl(e.target.value)}
                className="h-14 px-4 bg-white/40 border-white/50 rounded-2xl text-gray-900 placeholder:text-gray-500 focus:bg-white/60 focus:border-white/70 transition-all duration-200"
                required
              />
            </div>

            <div className="space-y-3">
              <label htmlFor="domain" className="block text-sm font-medium text-gray-800">
                Choose your domain
              </label>
              <Select value={domain} onValueChange={setDomain}>
                <SelectTrigger className="h-14 px-4 bg-white/40 border-white/50 rounded-2xl text-gray-900 focus:bg-white/60 focus:border-white/70 transition-all duration-200">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="pocketurl.zip">pocketurl.zip</SelectItem>
                  <SelectItem value="claimfreerobux.online">claimfreerobux.online</SelectItem>
                  <SelectItem value="freeiphoneskentucky.biz">freeiphoneskentucky.biz</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <Button 
              type="submit" 
              disabled={!url || isLoading}
              className="w-full h-14 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-2xl transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2 shadow-lg hover:shadow-xl"
            >
              {isLoading ? (
                <div className="w-5 h-5 border-2 border-white/20 border-t-white rounded-full animate-spin" />
              ) : (
                <>
                  Shorten URL
                  <ArrowRight className="w-5 h-5" />
                </>
              )}
            </Button>
          </form>

          {shortenedUrl && (
            <div className="mt-8 p-6 bg-white/30 border border-white/40 rounded-2xl backdrop-blur-sm">
              <p className="text-sm font-medium text-gray-800 mb-3">Your shortened URL:</p>
              <div className="flex gap-3">
                <Input
                  value={shortenedUrl}
                  readOnly
                  className="h-12 px-4 bg-white/60 border-white/60 rounded-xl text-gray-900 font-medium"
                />
                <Button
                  onClick={handleCopy}
                  className="h-12 px-6 bg-green-600 hover:bg-green-700 text-white rounded-xl transition-all duration-200 flex items-center gap-2"
                >
                  {copied ? (
                    <>
                      <CheckCircle className="w-4 h-4" />
                      Copied!
                    </>
                  ) : (
                    <>
                      <Copy className="w-4 h-4" />
                      Copy
                    </>
                  )}
                </Button>
              </div>
            </div>
          )}
        </div>
      </Card>

      {!shortenedUrl && (
        <div className="mt-12 text-center">
          <p className="text-gray-600 font-medium">
            No recent links
          </p>
        </div>
      )}
    </div>
  );
};

export default UrlShortener;