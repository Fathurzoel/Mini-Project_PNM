import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginForm from './components/LoginForm/LoginForm';
// import Home from './pages/Home/Home';
// import Quiz from './pages/Quiz/Quiz';
// import Results from './pages/Results/Results';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginForm />} />
        {/* <Route path="/Home" element={<Home />} /> */}
        {/* <Route path="/quiz" element={<Quiz />} />
        <Route path="/results" element={<Results />} /> */}
      </Routes>
    </Router>
  );
};

export default App;