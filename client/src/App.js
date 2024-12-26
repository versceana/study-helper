import './App.css';
import React from "react";
import { 
  BrowserRouter as Router, 
  Route, 
  Routes 
} from 'react-router-dom';
import Home from './pages/Home';
import Calendar from './pages/Calendar';
import Timer from './components/Timer';

function App() {
  return (
    <Router>
        <Routes>
          <Route path="/" element={<Timer />} />
          <Route path="/calendar" element={<Calendar />} />
        </Routes>
    </Router>
  );
}

export default App;
