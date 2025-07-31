import { useEffect, useState } from "react";
import { stockService } from "../services/stockService";
import { StockCard } from "../components/StockCard";
import type { StockResponse } from "../types/stock";

export function StockList() {
  const [stocks, setStocks] = useState<StockResponse[]>([]);
  const [filteredStocks, setFilteredStocks] = useState<StockResponse[]>([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedSector, setSelectedSector] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  // Pagination state
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [totalStocks, setTotalStocks] = useState(0);
  const [limit, setLimit] = useState(20);

  useEffect(() => {
    const handler = setTimeout(() => {
      if (searchTerm) {
        searchStocks();
      } else {
        loadStocks(currentPage, limit);
      }
    }, 1000); // 1000ms debounce delay

    return () => {
      clearTimeout(handler);
    };
  }, [currentPage, limit, searchTerm]);

  useEffect(() => {
    let filtered = stocks;

    if (selectedSector) {
      filtered = filtered.filter(
        (stock) => stock.profiles?.[0]?.sector === selectedSector
      );
    }

    setFilteredStocks(filtered);
  }, [selectedSector, stocks]);

  const loadStocks = async (page: number, itemsPerPage: number) => {
    try {
      setLoading(true);
      const response = await stockService.getAllStocks(page, itemsPerPage);

      // Handle the actual API response structure: { code: 200, message: "Success", data: { data: [...], total: ..., ... } }
      const apiResponse = response as any;

      if (apiResponse.data && apiResponse.data.data) {
        // Standard wrapped response
        setStocks(apiResponse.data.data || []);
        setFilteredStocks(apiResponse.data.data || []);
        setTotalPages(apiResponse.data.total_pages || 0);
        setTotalStocks(apiResponse.data.total || 0);
      } else {
        // Fallback - handle unexpected response format
        console.error("Unexpected response format:", response);
        setStocks([]);
        setFilteredStocks([]);
        setTotalPages(0);
        setTotalStocks(0);
      }

      if (!apiResponse.data?.data || apiResponse.data.data.length === 0) {
        console.log("No stocks found - backend may need data initialization");
      }
    } catch (err) {
      console.error("Error loading stocks:", err);
      setError(err instanceof Error ? err.message : "Failed to load stocks");
    } finally {
      setLoading(false);
    }
  };

  const searchStocks = async () => {
    try {
      setLoading(true);
      const response = await stockService.searchStocks(searchTerm);
      const apiResponse = response as any;

      if (apiResponse.data) {
        setStocks(apiResponse.data || []);
        setFilteredStocks(apiResponse.data || []);
        setTotalPages(1); // Reset pagination for search results
        setTotalStocks(apiResponse.data.length);
      } else {
        console.error("Unexpected search response format:", response);
        setStocks([]);
        setFilteredStocks([]);
      }
    } catch (err) {
      console.error("Error searching stocks:", err);
      setError(err instanceof Error ? err.message : "Failed to search stocks");
    } finally {
      setLoading(false);
    }
  };

  const handlePageChange = (newPage: number) => {
    if (newPage >= 1 && newPage <= totalPages) {
      setCurrentPage(newPage);
      window.scrollTo({ top: 0, behavior: "smooth" });
    }
  };

  const handleLimitChange = (newLimit: number) => {
    setLimit(newLimit);
    setCurrentPage(1); // Reset to first page when changing limit
  };

  const sectors = [
    ...new Set(
      (stocks || []).map((stock) => stock.profiles?.[0]?.sector).filter(Boolean)
    ),
  ];

  const renderPagination = () => {
    const pages = [];
    const maxVisiblePages = 5;
    let startPage = Math.max(1, currentPage - Math.floor(maxVisiblePages / 2));
    let endPage = Math.min(totalPages, startPage + maxVisiblePages - 1);

    if (endPage - startPage < maxVisiblePages - 1) {
      startPage = Math.max(1, endPage - maxVisiblePages + 1);
    }

    for (let i = startPage; i <= endPage; i++) {
      pages.push(
        <button
          key={i}
          onClick={() => handlePageChange(i)}
          className={`px-3 py-1 mx-1 rounded ${
            i === currentPage
              ? "bg-blue-600 text-white"
              : "bg-gray-200 text-gray-700 hover:bg-gray-300"
          }`}
        >
          {i}
        </button>
      );
    }

    return pages;
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-16 w-16 border-b-2 border-blue-600 mx-auto mb-4"></div>
          <p className="text-gray-600">Loading stocks...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            <p className="font-bold">Error</p>
            <p>{error}</p>
          </div>
          <button
            onClick={() => loadStocks(currentPage, limit)}
            className="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="bg-white/80 backdrop-blur-md shadow-sm">
        <div className="max-w-4xl mx-auto px-3 py-3">
          <div className="text-center mb-6">
            <h1 className="text-3xl font-bold text-gray-800">
              📈 Stock Market
            </h1>
            <p className="text-sm text-gray-500">
              Search and filter Indonesian listed companies
            </p>
          </div>

          <div className="max-w-3xl mx-auto">
            <div className="bg-gray-50 border rounded-xl p-4 shadow-sm mb-4">
              <div className="grid grid-cols-1 md:grid-cols-5 gap-3 items-end">
                {/* Search input */}
                <div className="md:col-span-3">
                  <label className="block text-xs font-semibold text-gray-600 mb-1">
                    🔍 Stock Name / Code
                  </label>
                  <input
                    type="text"
                    placeholder="e.g. TLKM, BBRI, etc."
                    value={searchTerm}
                    onChange={(e) => {
                      setSearchTerm(e.target.value);
                      setCurrentPage(1);
                    }}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:outline-none"
                  />
                </div>

                {/* Sector filter */}
                <div className="md:col-span-2">
                  <label className="block text-xs font-semibold text-gray-600 mb-1">
                    🏷️ Sector
                  </label>
                  <select
                    value={selectedSector}
                    onChange={(e) => setSelectedSector(e.target.value)}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:outline-none"
                  >
                    <option value="">All Sectors</option>
                    {sectors.map((sector) => (
                      <option key={sector} value={sector}>
                        {sector}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              {(searchTerm || selectedSector) && (
                <div className="mt-3 flex justify-between text-xs text-gray-600">
                  <span>
                    Showing {filteredStocks.length} of {totalStocks} stocks
                  </span>
                  <button
                    onClick={() => {
                      setSearchTerm("");
                      setSelectedSector("");
                    }}
                    className="text-blue-600 hover:underline"
                  >
                    Clear Filters
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-4 py-8">
        {filteredStocks.length > 0 ? (
          <>
            <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 px-4 sm:px-6 lg:px-8 mb-8">
              {filteredStocks.map((stock) => (
                <StockCard key={stock.code} stock={stock} />
              ))}
            </div>

            {/* Pagination Controls */}
            {totalPages > 1 && (
              <div className="flex flex-col items-center space-y-4">
                <div className="flex items-center space-x-2">
                  <button
                    onClick={() => handlePageChange(currentPage - 1)}
                    disabled={currentPage === 1}
                    className="px-3 py-1 rounded bg-gray-200 text-gray-700 hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    Previous
                  </button>

                  {renderPagination()}

                  <button
                    onClick={() => handlePageChange(currentPage + 1)}
                    disabled={currentPage === totalPages}
                    className="px-3 py-1 rounded bg-gray-200 text-gray-700 hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed"
                  >
                    Next
                  </button>
                </div>

                <div className="flex items-center space-x-4">
                  <span className="text-sm text-gray-600">
                    Page {currentPage} of {totalPages}
                  </span>

                  <select
                    value={limit}
                    onChange={(e) => handleLimitChange(Number(e.target.value))}
                    className="px-2 py-1 border border-gray-300 rounded text-sm"
                  >
                    <option value={10}>10 per page</option>
                    <option value={20}>20 per page</option>
                    <option value={50}>50 per page</option>
                    <option value={100}>100 per page</option>
                  </select>
                </div>
              </div>
            )}
          </>
        ) : (
          <div className="text-center py-16">
            <div className="text-6xl mb-4">📊</div>
            <h3 className="text-xl font-semibold text-gray-900 mb-2">
              No Stocks Available
            </h3>
            <p className="text-gray-600 mb-4">
              {searchTerm || selectedSector
                ? "Try adjusting your search criteria"
                : "The stock database appears to be empty"}
            </p>
            {!searchTerm && !selectedSector && (
              <div className="text-sm text-gray-500">
                <p className="mb-2">To populate stock data:</p>
                <ol className="text-left max-w-md mx-auto space-y-1">
                  <li>1. Ensure the Go backend is running on localhost:3000</li>
                  <li>2. Run the data import/sync process</li>
                  <li>3. Check the backend logs for any connection issues</li>
                </ol>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
