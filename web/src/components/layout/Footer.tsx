import { Link } from 'react-router-dom';

export function Footer() {
    return (
        <footer className="bg-gray-800 text-white mt-auto">
            <div className="container mx-auto px-4 py-8">
                <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                    <div>
                        <h3 className="text-lg font-semibold mb-4">Go Stock</h3>
                        <p className="text-gray-300 text-sm">
                            Comprehensive stock market data and analysis platform built with modern web technologies.
                        </p>
                    </div>

                    <div>
                        <h3 className="text-lg font-semibold mb-4">Quick Links</h3>
                        <ul className="space-y-2 text-sm">
                            <li>
                                <Link to="/stocks" className="text-gray-300 hover:text-white transition-colors">
                                    Browse Stocks
                                </Link>
                            </li>
                            <li>
                                <Link to="/" className="text-gray-300 hover:text-white transition-colors">
                                    Home
                                </Link>
                            </li>
                        </ul>
                    </div>

                    <div>
                        <h3 className="text-lg font-semibold mb-4">Data Sources</h3>
                        <p className="text-gray-300 text-sm">
                            Data fetched from IDX and comprehensive financial reports.
                        </p>
                    </div>
                </div>

                <div className="border-t border-gray-700 mt-8 pt-8 text-center">
                    <p className="text-gray-400 text-sm">
                        © {new Date().getFullYear()} Go Stock. Built with React, TypeScript, and Tailwind CSS.
                    </p>
                </div>
            </div>
        </footer>
    );
}
