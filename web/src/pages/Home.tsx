import { Link } from 'react-router-dom';

export default function Home() {
    return (
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
            <div className="container mx-auto px-4 py-16">
                <div className="text-center mb-16">
                    <h1 className="text-5xl font-bold text-gray-900 mb-6">
                        Welcome to <span className="text-blue-600">Go Stock</span>
                    </h1>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        Comprehensive stock market data and analysis platform. 
                        Access stock information, financial reports, and broker trading activity.
                    </p>
                </div>

            
                <div className="text-center mt-16">
                    <Link 
                        to="/stocks" 
                        className="inline-block bg-blue-600 text-white px-8 py-4 rounded-lg text-lg font-semibold hover:bg-blue-700 transition-colors shadow-lg"
                    >
                        Get Started →
                    </Link>
                </div>
            </div>
        </div>
    );
}
