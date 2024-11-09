import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Auth } from "../pages/auth";
import { Register } from "../pages/register"
import { Main } from "../pages/main"
import { Money } from "../pages/money/Money";

const AppRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main />}/>
        <Route path="/auth" element={<Auth />} />
        <Route path="/register" element={<Register />} />
        <Route path="/money" element={<Money />}/>
      </Routes>
    </BrowserRouter>
  );
};

export { AppRoutes };