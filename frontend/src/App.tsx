import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import AutonomousMobileRobotRegisterCreate from "./components/AutonomousMobileRobotRegisterCreate";
export default function App() {
  return (
    <Router>
      <div>
        <Navbar />
        <Routes>
          <Route path="/" element={<AutonomousMobileRobotRegisterCreate />} />
        </Routes>
      </div>
    </Router>
  );
}