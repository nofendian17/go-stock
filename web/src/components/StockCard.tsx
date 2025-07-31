import React from "react";
import { useNavigate } from "react-router-dom";
import type { StockResponse } from "../types/stock";

interface StockCardProps {
  stock: StockResponse;
}

export const StockCard: React.FC<StockCardProps> = ({ stock }) => {
  const navigate = useNavigate();
  const profile = stock.profiles?.[0];

  const handleClick = () => {
    navigate(`/stocks/${stock.code}`);
  };

  const formatDate = (date?: string) =>
    date ? new Date(date).toLocaleDateString("id-ID") : "-";

  const formatNumber = (val?: number) =>
    typeof val === "number" ? val.toLocaleString("id-ID") : "-";

  return (
    <div
      onClick={handleClick}
      className="cursor-pointer rounded-xl border border-gray-200 bg-white shadow-sm transition hover:shadow-md"
    >
      <div className="p-4 space-y-4">
        <div className="flex items-start justify-between">
          <div>
            <h3 className="text-lg font-semibold text-gray-900">{stock.name}</h3>
            <p className="text-xs text-blue-600 font-mono font-medium">{stock.code}</p>
          </div>
          {profile?.logo && (
            <div className="w-12 h-12 shrink-0 rounded-md bg-gray-100 overflow-hidden flex items-center justify-center">
              <img
                src={profile.logo}
                alt={stock.name}
                className="w-full h-full object-cover"
                width={48}
                height={48}
                loading="lazy"
                decoding="async"
                referrerPolicy="no-referrer"
                onError={(e) => {
                  e.currentTarget.src = `https://placehold.co/48x48?text=${stock.code}`;
                }}
              />
            </div>
          )}
        </div>

        {profile && (
          <div className="text-sm space-y-2">
            <div className="flex justify-between">
              <span className="text-gray-500">Sector</span>
              <span className="font-medium text-gray-800">{profile.sector || "-"}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-500">Listed</span>
              <span className="font-medium text-gray-800">{formatDate(profile.listing_date)}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-500">Shares</span>
              <span className="font-medium text-gray-800">{formatNumber(stock.share)}</span>
            </div>
          </div>
        )}

        {profile?.main_business && (
          <div className="pt-3 border-t border-gray-100">
            <p className="text-sm text-gray-600 line-clamp-2">{profile.main_business}</p>
          </div>
        )}
      </div>

      <div className="bg-gray-50 px-4 py-2 text-xs text-gray-500 flex justify-between items-center">
        <span>Click to view details</span>
        <span className="text-blue-600 font-medium">→</span>
      </div>
    </div>
  );
};
