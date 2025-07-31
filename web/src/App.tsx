// src/App.tsx
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { MainLayout } from "./components/layout/MainLayout";
import Home from "./pages/Home";
import { StockList } from "./pages/StockList";
import { StockDetail } from "./pages/StockDetail";

function App() {
    return (
        <BrowserRouter>
            <MainLayout>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/stocks" element={<StockList />} />
                    <Route path="/stocks/:stockCode" element={<StockDetail />} />
                </Routes>
            </MainLayout>
        </BrowserRouter>
    );
}

export default App;
