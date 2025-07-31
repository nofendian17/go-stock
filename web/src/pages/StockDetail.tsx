import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { stockService } from "../services/stockService";
import { StockChart } from "../components/StockChart";
import { FinancialReportCard } from "../components/FinancialReportCard";
import type {
  StockResponse,
  StockSummaryResponse,
  FinancialReportResponse,
  BrokerSummaryResponse,
} from "../types/stock";

export function StockDetail() {
  const { stockCode } = useParams<{ stockCode: string }>();
  const [stock, setStock] = useState<StockResponse | null>(null);
  const [summaries, setSummaries] = useState<StockSummaryResponse[]>([]);
  const [reports, setReports] = useState<FinancialReportResponse[]>([]);
  const [brokerSummary, setBrokerSummary] =
    useState<BrokerSummaryResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [dateRange, setDateRange] = useState({
    startDate: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000)
      .toISOString()
      .split("T")[0],
    endDate: new Date().toISOString().split("T")[0],
  });
  const [brokerParams, setBrokerParams] = useState({
    investorType: "ALL",
    transactionType: "ALL",
  });
  const [activeTab, setActiveTab] = useState("company");

  useEffect(() => {
    if (stockCode) {
      loadStockData();
      loadFinancialReports();
      loadBrokerSummary();
    }
  }, [stockCode]);

  const loadStockData = async () => {
    try {
      setLoading(true);
      const [stockData, summaryData] = await Promise.all([
        stockService.getStockByCode(stockCode!),
        stockService.getStockSummaries(
          stockCode!,
          dateRange.startDate,
          dateRange.endDate
        ),
      ]);

      // Handle wrapped response structure
      const stockResponse = stockData as any;
      const summariesResponse = summaryData as any;

      setStock(stockResponse.data || stockResponse);
      setSummaries(
        Array.isArray(summariesResponse.data)
          ? summariesResponse.data
          : Array.isArray(summariesResponse)
          ? summariesResponse
          : []
      );
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Failed to load stock data"
      );
    } finally {
      setLoading(false);
    }
  };

  const loadFinancialReports = async () => {
    if (!stockCode) return;

    try {
      // Load reports for different periods for the last 5 years
      const periods = ["TW1", "TW2", "TW3", "Audit"];
      const currentYear = new Date().getFullYear();
      const years = Array.from({ length: 5 }, (_, i) =>
        (currentYear - i).toString()
      );
      const reportPromises = [];

      for (const year of years) {
        for (const period of periods) {
          reportPromises.push(
            stockService
              .getFinancialReport(stockCode, period, year)
              .then((response) => {
                const wrappedResponse = response as any;
                return wrappedResponse.data || response;
              })
              .catch(() => null)
          );
        }
      }

      const reports = await Promise.all(reportPromises);
      setReports(reports.filter(Boolean) as FinancialReportResponse[]);
    } catch (err) {
      console.error("Failed to load financial reports:", err);
    }
  };

  const loadBrokerSummary = async () => {
    if (!stockCode || !dateRange.startDate || !dateRange.endDate) return;

    try {
      const data = await stockService.getBrokerSummaries(
        stockCode,
        dateRange.startDate,
        dateRange.endDate,
        brokerParams.investorType,
        brokerParams.transactionType
      );
      

      // Handle wrapped response structure
      const response = data as any;
      setBrokerSummary(response.data || response);
    } catch (err) {
      console.error("Failed to load broker summary:", err);
      setBrokerSummary(null);
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center py-8">
        <p className="text-red-600 mb-4">Error: {error}</p>
        <button
          onClick={loadStockData}
          className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
        >
          Retry
        </button>
      </div>
    );
  }

  if (!stock) {
    return <div className="text-center py-8">Stock not found</div>;
  }

  const profile = stock.profiles?.[0];

  return (
    <div className="container mx-auto px-4 py-4">
      <div className="bg-white rounded-lg shadow p-4 mb-4">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold text-gray-800">{stock.name}</h1>
            <div className="flex items-center space-x-2 mt-1">
              <span className="text-md font-mono font-semibold text-blue-600 bg-blue-50 px-2 py-0.5 rounded">
                {stock.code}
              </span>
              <span className="text-xs text-gray-500">• {stock.board}</span>
            </div>
          </div>
          {profile?.logo && (
            <img
              src={profile.logo}
              alt={stock.name}
              className="w-16 h-16 rounded-md object-contain border border-gray-200"
            />
          )}
        </div>
      </div>

      <div className="bg-white rounded-lg shadow p-4 mb-4">
        <div className="mb-4">
          <div className="flex items-center space-x-2">
            <input
              type="date"
              value={dateRange.startDate}
              onChange={(e) =>
                setDateRange({ ...dateRange, startDate: e.target.value })
              }
              className="border rounded px-2 py-1 text-sm"
            />
            <input
              type="date"
              value={dateRange.endDate}
              onChange={(e) =>
                setDateRange({ ...dateRange, endDate: e.target.value })
              }
              className="border rounded px-2 py-1 text-sm"
            />
            <button
              onClick={loadStockData}
              className="px-3 py-1 bg-blue-600 text-white rounded hover:bg-blue-700 text-sm"
            >
              Update
            </button>
          </div>
        </div>
        <StockChart summaries={summaries} title="" />
      </div>

      {/* Tab Navigation */}
      <div className="mb-4 border-b border-gray-200">
        <nav className="flex space-x-8">
          <button
            onClick={() => setActiveTab("company")}
            className={`py-4 px-1 border-b-2 font-medium text-sm ${
              activeTab === "company"
                ? "border-blue-500 text-blue-600"
                : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
            }`}
          >
            Company Information
          </button>
          <button
            onClick={() => setActiveTab("financial")}
            className={`py-4 px-1 border-b-2 font-medium text-sm ${
              activeTab === "financial"
                ? "border-blue-500 text-blue-600"
                : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
            }`}
          >
            Financial Reports
          </button>
          <button
            onClick={() => setActiveTab("broker")}
            className={`py-4 px-1 border-b-2 font-medium text-sm ${
              activeTab === "broker"
                ? "border-blue-500 text-blue-600"
                : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
            }`}
          >
            Broker Summary
          </button>
        </nav>
      </div>

      {/* Tab Content */}
      <div className="bg-white rounded-lg shadow p-4">
        {activeTab === "company" && (
          <div className="space-y-6">
            {/* Company Overview Card */}
            <div className="bg-white rounded-lg shadow p-6">
              <div className="flex flex-col md:flex-row gap-6">
                {profile?.logo && (
                  <div className="flex-shrink-0">
                    <img
                      src={profile.logo}
                      alt={stock.name}
                      className="w-32 h-32 rounded-md object-contain border border-gray-200"
                    />
                  </div>
                )}
                <div className="flex-1">
                  <h2 className="text-xl font-bold text-gray-800">
                    {stock.name}
                  </h2>
                  <div className="flex items-center gap-2 mt-1 mb-4">
                    <span className="text-sm font-mono font-semibold text-blue-600 bg-blue-50 px-2 py-0.5 rounded">
                      {stock.code}
                    </span>
                    <span className="text-xs text-gray-500">
                      • {stock.board}
                    </span>
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <h3 className="text-sm font-medium text-gray-500 mb-1">
                        Sector
                      </h3>
                      <p className="text-gray-800">
                        {profile?.sector} ({profile?.sub_sector})
                      </p>
                    </div>
                    <div>
                      <h3 className="text-sm font-medium text-gray-500 mb-1">
                        Industry
                      </h3>
                      <p className="text-gray-800">
                        {profile?.industry} ({profile?.sub_industry})
                      </p>
                    </div>
                    <div>
                      <h3 className="text-sm font-medium text-gray-500 mb-1">
                        Listed Date
                      </h3>
                      <p className="text-gray-800">
                        {new Date(
                          profile?.listing_date || stock.listing_date
                        ).toLocaleDateString()}
                      </p>
                    </div>
                    <div>
                      <h3 className="text-sm font-medium text-gray-500 mb-1">
                        Shares Outstanding
                      </h3>
                      <p className="text-gray-800">
                        {stock.share?.toLocaleString()}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            {/* Contact Information Card */}
            <div className="bg-white rounded-lg shadow p-6">
              <h3 className="text-lg font-semibold mb-4">
                Contact Information
              </h3>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">
                    Address
                  </h4>
                  <p className="text-gray-800 whitespace-pre-line">
                    {profile?.address}
                  </p>
                </div>
                <div className="space-y-4">
                  <div>
                    <h4 className="text-sm font-medium text-gray-500 mb-1">
                      Phone
                    </h4>
                    <p className="text-gray-800">{profile?.phone}</p>
                  </div>
                  <div>
                    <h4 className="text-sm font-medium text-gray-500 mb-1">
                      Fax
                    </h4>
                    <p className="text-gray-800">{profile?.fax}</p>
                  </div>
                  <div>
                    <h4 className="text-sm font-medium text-gray-500 mb-1">
                      Email
                    </h4>
                    <a
                      href={`mailto:${profile?.email}`}
                      className="text-blue-600 hover:underline"
                    >
                      {profile?.email}
                    </a>
                  </div>
                  <div>
                    <h4 className="text-sm font-medium text-gray-500 mb-1">
                      Website
                    </h4>
                    <a
                      href={
                        profile?.website?.startsWith("http")
                          ? profile.website
                          : `https://${profile?.website}`
                      }
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-blue-600 hover:underline"
                    >
                      {profile?.website}
                    </a>
                  </div>
                </div>
              </div>
            </div>

            {/* Business Information Card */}
            <div className="bg-white rounded-lg shadow p-6">
              <h3 className="text-lg font-semibold mb-4">
                Business Information
              </h3>
              <div>
                <h4 className="text-sm font-medium text-gray-500 mb-1">
                  Main Business
                </h4>
                <p className="text-gray-800">{profile?.main_business}</p>
              </div>
            </div>

            {/* Corporate Structure */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              {/* Directors */}
              <div className="bg-white rounded-lg shadow p-6">
                <h3 className="text-lg font-semibold mb-4">Directors</h3>
                <div className="space-y-3">
                  {stock.directors?.map((director, index) => (
                    <div
                      key={index}
                      className="border-b pb-2 last:border-0 last:pb-0"
                    >
                      <p className="font-medium text-gray-800">
                        {director.name}
                      </p>
                      <p className="text-sm text-gray-600">
                        {director.position}
                      </p>
                    </div>
                  ))}
                </div>
              </div>

              {/* Commissioners */}
              <div className="bg-white rounded-lg shadow p-6">
                <h3 className="text-lg font-semibold mb-4">Commissioners</h3>
                <div className="space-y-3">
                  {stock.commissioners?.map((commissioner, index) => (
                    <div
                      key={index}
                      className="border-b pb-2 last:border-0 last:pb-0"
                    >
                      <p className="font-medium text-gray-800">
                        {commissioner.name}
                      </p>
                      <p className="text-sm text-gray-600">
                        {commissioner.position}
                        {commissioner.is_independent && (
                          <span className="ml-2 text-xs bg-green-100 text-green-800 px-2 py-0.5 rounded">
                            Independent
                          </span>
                        )}
                      </p>
                    </div>
                  ))}
                </div>
              </div>

              {/* Audit Committee */}
              <div className="bg-white rounded-lg shadow p-6">
                <h3 className="text-lg font-semibold mb-4">Audit Committee</h3>
                <div className="space-y-3">
                  {stock.audit_committees?.map((member, index) => (
                    <div
                      key={index}
                      className="border-b pb-2 last:border-0 last:pb-0"
                    >
                      <p className="font-medium text-gray-800">{member.name}</p>
                      <p className="text-sm text-gray-600">{member.position}</p>
                    </div>
                  ))}
                </div>
              </div>
            </div>

            {/* Major Shareholders */}
            {stock.shareholders && stock.shareholders.length > 0 && (
              <div className="bg-white rounded-lg shadow p-6">
                <h3 className="text-lg font-semibold mb-4">
                  Major Shareholders
                </h3>
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Name
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Shares
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Percentage
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Category
                        </th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {stock.shareholders
                        .filter((sh) => sh.percentage > 0)
                        .sort((a, b) => b.percentage - a.percentage)
                        .map((shareholder, index) => (
                          <tr key={index}>
                            <td className="px-4 py-2 whitespace-nowrap text-sm text-gray-800">
                              {shareholder.name}
                              {shareholder.is_controller && (
                                <span className="ml-2 text-xs bg-yellow-100 text-yellow-800 px-1.5 py-0.5 rounded">
                                  Controller
                                </span>
                              )}
                            </td>
                            <td className="px-4 py-2 whitespace-nowrap text-sm text-gray-800">
                              {shareholder.share.toLocaleString()}
                            </td>
                            <td className="px-4 py-2 whitespace-nowrap text-sm text-gray-800">
                              {shareholder.percentage.toFixed(2)}%
                            </td>
                            <td className="px-4 py-2 whitespace-nowrap text-sm text-gray-800">
                              {shareholder.category}
                            </td>
                          </tr>
                        ))}
                    </tbody>
                  </table>
                </div>
              </div>
            )}

            {/* Subsidiaries */}
            {stock.subsidiaries && stock.subsidiaries.length > 0 && (
              <div className="bg-white rounded-lg shadow p-6">
                <h3 className="text-lg font-semibold mb-4">Subsidiaries</h3>
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Name
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Business
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Location
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Ownership
                        </th>
                        <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Status
                        </th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {stock.subsidiaries
                        .filter(
                          (sub, index, self) =>
                            index ===
                            self.findIndex(
                              (s) =>
                                s.name === sub.name &&
                                s.location === sub.location
                            )
                        )
                        .map((subsidiary, index) => (
                          <tr key={index}>
                            <td className="px-4 py-2 text-sm text-gray-800">
                              {subsidiary.name}
                            </td>
                            <td className="px-4 py-2 text-sm text-gray-800">
                              {subsidiary.business_fields}
                            </td>
                            <td className="px-4 py-2 text-sm text-gray-800">
                              {subsidiary.location}
                            </td>
                            <td className="px-4 py-2 text-sm text-gray-800">
                              {subsidiary.percentage}%
                            </td>
                            <td className="px-4 py-2 whitespace-nowrap">
                              <span
                                className={`px-2 py-1 text-xs rounded-full ${
                                  subsidiary.operation_status === "Aktif"
                                    ? "bg-green-100 text-green-800"
                                    : "bg-gray-100 text-gray-800"
                                }`}
                              >
                                {subsidiary.operation_status}
                              </span>
                            </td>
                          </tr>
                        ))}
                    </tbody>
                  </table>
                </div>
              </div>
            )}
          </div>
        )}

        {activeTab === "financial" && (
          <div>
            <h3 className="text-md font-semibold mb-3">Financial Reports</h3>
            <div className="space-y-3">
              {reports.length > 0 ? (
                reports
                  .slice(0, 5)
                  .map((report) => (
                    <FinancialReportCard
                      key={`${report.report_period}-${report.report_year}`}
                      report={report}
                    />
                  ))
              ) : (
                <p className="text-gray-500 text-sm">
                  No financial reports available
                </p>
              )}
            </div>
          </div>
        )}

        {activeTab === "broker" && (
          <div className="space-y-4">
            {/* Filter Controls - 4 Column Layout */}
            <div className="bg-gray-50 p-4 rounded-lg">
              <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
                {/* Start Date */}
                <div>
                  <label className="block text-xs font-medium text-gray-600 mb-1">
                    Start Date
                  </label>
                  <input
                    type="date"
                    value={dateRange.startDate}
                    onChange={(e) =>
                      setDateRange({ ...dateRange, startDate: e.target.value })
                    }
                    className="w-full border rounded px-3 py-2 text-sm focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>

                {/* End Date */}
                <div>
                  <label className="block text-xs font-medium text-gray-600 mb-1">
                    End Date
                  </label>
                  <input
                    type="date"
                    value={dateRange.endDate}
                    onChange={(e) =>
                      setDateRange({ ...dateRange, endDate: e.target.value })
                    }
                    className="w-full border rounded px-3 py-2 text-sm focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>

                {/* Investor Type */}
                <div>
                  <label className="block text-xs font-medium text-gray-600 mb-1">
                    Investor Type
                  </label>
                  <select
                    value={brokerParams.investorType}
                    onChange={(e) =>
                      setBrokerParams({
                        ...brokerParams,
                        investorType: e.target.value,
                      })
                    }
                    className="w-full border rounded px-3 py-2 text-sm focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="ALL">All Investors</option>
                    <option value="D">Domestic</option>
                    <option value="F">Foreign</option>
                  </select>
                </div>

                {/* Transaction Type */}
                <div>
                  <label className="block text-xs font-medium text-gray-600 mb-1">
                    Transaction Type
                  </label>
                  <select
                    value={brokerParams.transactionType}
                    onChange={(e) =>
                      setBrokerParams({
                        ...brokerParams,
                        transactionType: e.target.value,
                      })
                    }
                    className="w-full border rounded px-3 py-2 text-sm focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="ALL">All Types</option>
                    <option value="RG">Regular</option>
                    <option value="TN">Tunai</option>
                    <option value="NG">Nego</option>
                  </select>
                </div>
              </div>

              <div className="mt-4 flex justify-end">
                <button
                  onClick={loadBrokerSummary}
                  className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 text-sm font-medium"
                >
                  Apply Filters
                </button>
              </div>
            </div>

            {/* Broker Summary Table */}
            <div className="bg-white rounded-lg shadow overflow-hidden">
              {brokerSummary ? (
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th
                          colSpan={4}
                          className="px-6 py-5 text-left text-xs font-medium text-gray-500 uppercase tracking-wider bg-gray-100"
                        >
                          Buyers
                        </th>
                        <th
                          colSpan={4}
                          className="px-6 py-5 text-left text-xs font-medium text-gray-500 uppercase tracking-wider bg-gray-100"
                        >
                          Sellers
                        </th>
                      </tr>
                      <tr>
                        {/* Buyer Columns */}
                        {["Broker", "Lot", "Value", "Average"].map(
                          (title, idx) => (
                            <th
                              key={`buyer-header-${idx}`}
                              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                            >
                              {title}
                            </th>
                          )
                        )}
                        {/* Seller Columns */}
                        {["Broker", "Lot", "Value", "Average"].map(
                          (title, idx) => (
                            <th
                              key={`seller-header-${idx}`}
                              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                            >
                              {title}
                            </th>
                          )
                        )}
                      </tr>
                    </thead>

                    <tbody className="bg-white divide-y divide-gray-200">
                      {brokerSummary.buyers?.map((buyer, index) => {
                        const seller = brokerSummary.sellers?.[index] || {};
                        return (
                          <tr
                            key={index}
                            className={
                              index % 2 === 0 ? "bg-white" : "bg-gray-50"
                            }
                          >
                            {/* Buyer Data */}
                            <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                              {buyer.broker_code}
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {buyer.lot.toLocaleString()}
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {buyer.val.toLocaleString()}
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {buyer.avg.toLocaleString()}
                            </td>

                            {/* Seller Data */}
                            <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                              {seller.broker_code || "-"}
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {seller.lot ? seller.lot.toLocaleString() : "-"}
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {seller.val ? seller.val.toLocaleString() : "-"}
                            </td>
                            <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                              {seller.avg ? seller.avg.toLocaleString() : "-"}
                            </td>
                          </tr>
                        );
                      })}
                    </tbody>

                    <tfoot className="bg-gray-50">
                      <tr>
                        <th
                          colSpan={8}
                          className="px-6 py-5 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
                        >
                          Summary
                          <div className="mt-2 font-normal text-sm text-gray-700 space-x-4">
                            <span>
                              Total Lot:{" "}
                              {brokerSummary.summary.total_lot.toLocaleString()}
                            </span>
                            <span>
                              Total Value:{" "}
                              {brokerSummary.summary.total_val.toLocaleString()}
                            </span>
                            <span>
                              Foreign Net:{" "}
                              {brokerSummary.summary.foreign_net_val.toLocaleString()}
                            </span>
                            <span>
                              Average:{" "}
                              {brokerSummary.summary.avg.toLocaleString()}
                            </span>
                          </div>
                        </th>
                      </tr>
                    </tfoot>
                  </table>
                </div>
              ) : (
                <div className="p-8 text-center text-gray-500">
                  No broker data available for selected filters
                </div>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}