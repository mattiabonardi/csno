import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import { TRexGamePage } from "./pages/games/t-rex-game/[id]";
import Home from "./pages/Home";
import NoPage from "./pages/NoPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" index element={<Home />} />
        <Route path="/games/t-rex-game/:id" element={<TRexGamePage />} />
        <Route path="*" element={<NoPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
