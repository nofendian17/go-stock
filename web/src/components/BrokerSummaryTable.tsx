import React from 'react';
import type { BrokerSummaryResponse } from '../types/stock';

interface BrokerSummaryTableProps {
    data: BrokerSummaryResponse;
}

export const BrokerSummaryTable: React.FC<BrokerSummaryTableProps> = ({ data }) => {
    // Display values exactly as received from API
    const displayValue = (value: string | number) => String(value);

    // Ensure data properties exist and are arrays
    const buyers = Array.isArray(data?.buyers) ? data.buyers : [];
    const sellers = Array.isArray(data?.sellers) ? data.sellers : [];
    const summary = data?.summary || {
        total_lot: 0,
        total_val: '0',
        avg: 0,
        foreign_net_val: '0'
    };

    return (
        <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-lg font-semibold mb-4">
                Broker Summary ({data?.start_date || 'N/A'} - {data?.end_date || 'N/A'})
            </h3>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <h4 className="text-md font-medium text-green-600 mb-3">Top Buyers</h4>
                    <div className="overflow-x-auto">
                        <table className="min-w-full divide-y divide-gray-200">
                            <thead>
                                <tr>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Broker</th>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Lot</th>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Value</th>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Avg</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-200">
                                {buyers.slice(0, 10).map((buyer) => (
                                    <tr key={buyer.broker_code}>
                                        <td className="px-4 py-2 text-sm text-gray-900 font-medium">{buyer.broker_code}</td>
                                        <td className="px-4 py-2 text-sm text-gray-500">{displayValue(buyer.lot)}</td>
                                        <td className="px-4 py-2 text-sm text-gray-500 font-medium">{displayValue(buyer.val)}</td>
                                        <td className="px-4 py-2 text-sm text-gray-500">{displayValue(buyer.avg)}</td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </div>

                <div>
                    <h4 className="text-md font-medium text-red-600 mb-3">Top Sellers</h4>
                    <div className="overflow-x-auto">
                        <table className="min-w-full divide-y divide-gray-200">
                            <thead>
                                <tr>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Broker</th>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Lot</th>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Value</th>
                                    <th className="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Avg</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-200">
                                {sellers.slice(0, 10).map((seller) => (
                                    <tr key={seller.broker_code}>
                                        <td className="px-4 py-2 text-sm text-gray-900 font-medium">{seller.broker_code}</td>
                                        <td className="px-4 py-2 text-sm text-gray-500">{displayValue(seller.lot)}</td>
                                        <td className="px-4 py-2 text-sm text-gray-500 font-medium">{displayValue(seller.val)}</td>
                                        <td className="px-4 py-2 text-sm text-gray-500">{displayValue(seller.avg)}</td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>

            <div className="mt-6 pt-4 border-t">
                <h4 className="text-md font-medium mb-2">Summary</h4>
                <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
                    <div className="bg-gray-50 p-3 rounded">
                        <p className="text-sm text-gray-600">Total Lot</p>
                        <p className="text-lg font-semibold">{displayValue(summary.total_lot)}</p>
                    </div>
                    <div className="bg-gray-50 p-3 rounded">
                        <p className="text-sm text-gray-600">Total Value</p>
                        <p className="text-lg font-semibold">{displayValue(summary.total_val)}</p>
                    </div>
                    <div className="bg-gray-50 p-3 rounded">
                        <p className="text-sm text-gray-600">Average</p>
                        <p className="text-lg font-semibold">{displayValue(summary.avg)}</p>
                    </div>
                    <div className="bg-gray-50 p-3 rounded">
                        <p className="text-sm text-gray-600">Foreign Net</p>
                        <p className="text-lg font-semibold">{displayValue(summary.foreign_net_val)}</p>
                    </div>
                </div>
            </div>
        </div>
    );
};
