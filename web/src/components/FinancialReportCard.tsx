import React from 'react';
import type { FinancialReportResponse } from '../types/stock';

interface FinancialReportCardProps {
    report: FinancialReportResponse;
}

export const FinancialReportCard: React.FC<FinancialReportCardProps> = ({ report }) => {
    return (
        <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
                <div>
                    <h3 className="text-lg font-bold text-gray-900">
                        {report.stock_name} ({report.stock_code})
                    </h3>
                    <p className="text-sm text-gray-600">
                        {report.report_period} - {report.report_year}
                    </p>
                </div>
                <div className="text-right">
                    <p className="text-xs text-gray-500">
                        Updated: {new Date(report.file_modified).toLocaleDateString()}
                    </p>
                </div>
            </div>

            {report.attachment && report.attachment.length > 0 && (
                <div className="mt-4">
                    <h4 className="text-sm font-semibold text-gray-700 mb-2">Attachments:</h4>
                    <div className="space-y-2">
                        {report.attachment.map((file) => (
                            <div key={file.file_id} className="flex items-center justify-between p-2 bg-gray-50 rounded">
                                <div>
                                    <p className="text-sm font-medium text-gray-900">{file.file_name}</p>
                                    <p className="text-xs text-gray-500">
                                        {(file.file_size / 1024 / 1024).toFixed(2)} MB
                                    </p>
                                </div>
                                <a
                                    href={file.file_path}
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    className="text-blue-600 hover:text-blue-800 text-sm font-medium"
                                >
                                    Download
                                </a>
                            </div>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};
