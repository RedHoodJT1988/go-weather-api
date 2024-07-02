import React, { useState } from "react";

const Weather: React.FC = () => {
  const [lat, setLat] = useState('');
  const [long, setLong] = useState('');
  const [weather, setWeather] = useState<{ condition: string, temperature: string} | null>(null);
  const [error, setError] = useState<string | null>(null);

  const fetchWeather = async () => {
    try {
      const response = await fetch(`/weather?lat=${lat}&long${long}`);
      if (!response.ok) {
        throw new Error('Failed to fetch weather data');
      }
      const data = await response.json();
      setWeather(data);
      setError(null);
    } catch (err: any) {
      setError(err.message);
      setWeather(null);
    }
  };

  return (
    <div>
      <h1>Weather Checker</h1>
      <input
        type="text"
        value={lat}
        onChange={(e) => setLat(e.target.value)}
        placeholder="Latitude"
      />
      <input
        type="text"
        value={long}
        onChange={(e) => setLong(e.target.value)}
        placeholder="Longitude"
      />
      <button onClick={fetchWeather}>Check Weather</button>
      {error && <p>Error: {error}</p>}
      {weather && (
        <div>
          <p>Condition: {weather.condition}</p>
          <p>Temperature: {weather.temperature}</p>
        </div>
      )}
    </div>
  );
};

export default Weather;