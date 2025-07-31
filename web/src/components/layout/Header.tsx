import { Link, NavLink } from 'react-router-dom';

export function Header() {
    const navLinkClasses = ({ isActive }: { isActive: boolean }) =>
        `px-3 py-2 rounded-md text-sm font-medium transition-colors ${
            isActive
                ? 'bg-blue-700 text-white'
                : 'text-blue-100 hover:bg-blue-700 hover:text-white'
        }`;

    return (
        <header className="bg-gradient-to-r from-blue-600 to-blue-800 text-white shadow-lg">
            <div className="container mx-auto px-4">
                <div className="flex items-center justify-between h-16">
                    <Link to="/" className="text-xl font-bold flex items-center space-x-2">
                        <span className="text-2xl">📈</span>
                        <span>Go Stock</span>
                    </Link>

                    <nav className="hidden md:flex space-x-8">
                        <NavLink to="/" className={navLinkClasses}>
                            Home
                        </NavLink>
                        <NavLink to="/stocks" className={navLinkClasses}>
                            Stocks
                        </NavLink>
                    </nav>

                    <div className="flex items-center space-x-4">
                        <div className="hidden md:block text-sm">
                            <span className="text-blue-200">Using EOD data</span>
                        </div>
                    </div>
                </div>
            </div>
        </header>
    );
}
