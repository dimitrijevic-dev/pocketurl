import Header from "@/components/Header";
import UrlShortener from "@/components/UrlShortener";
import Footer from "@/components/Footer";

const Index = () => {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      
      <main className="flex-1 flex flex-col items-center justify-center px-8 py-16">
        <div className="text-center mb-16 max-w-4xl">
          <h1 className="text-6xl md:text-7xl font-bold text-gray-900 mb-6 tracking-tight leading-none">
            The most versatile{" "}
            <span className="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
              URL shortener
            </span>
            , written in{" "}
            <span className="text-blue-600 font-semibold">Go</span>.
          </h1>
        </div>
        
        <UrlShortener />
      </main>
      
      <Footer />
    </div>
  );
};

export default Index;
