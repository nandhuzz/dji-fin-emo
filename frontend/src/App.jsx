
import { Routes, Route } from 'react-router-dom';
import Home from '@pages/Home';
import Login from '@pages/Login';
import HealthCheck from '@pages/HealthCheck'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/status" element={<HealthCheck/>} />
    </Routes>
  );
}

export default App;