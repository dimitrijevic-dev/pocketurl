import { Github, Linkedin } from "lucide-react";

const Header = () => {
  return (
    <header className="w-full backdrop-blur-xl bg-white/10 border-b border-white/10">
      <div className="flex justify-between items-center px-8 py-6 max-w-4xl mx-auto">
        <div className="flex items-center">
          <h1 className="text-2xl font-semibold text-gray-900 tracking-tight">
            PocketURL
          </h1>
        </div>
        
        <div className="flex items-center gap-6">
          <a 
            href="https://github.com/dimitrijevic-dev" 
            target="_blank" 
            rel="noopener noreferrer"
            className="text-gray-700 hover:text-gray-900 transition-colors duration-200 p-2 rounded-full hover:bg-white/20"
          >
            <Github size={22} />
          </a>
          <a 
            href="https://www.linkedin.com/in/dimitrijevic-dev/" 
            target="_blank" 
            rel="noopener noreferrer"
            className="text-gray-700 hover:text-gray-900 transition-colors duration-200 p-2 rounded-full hover:bg-white/20"
          >
            <Linkedin size={22} />
          </a>
        </div>
      </div>
    </header>
  );
};

export default Header;