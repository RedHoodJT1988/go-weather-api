import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import Weather from '../components/Weather';

test('renders Weather component and checks initial state', () => {
  render(<Weather />);
  expect(screen.getByPlaceholderText('Latitude')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Longitude')).toBeInTheDocument();
  expect(screen.getByText('Check Weather')).toBeInTheDocument();
});

test('fetches weather data on button click', async () => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      ok: true,
      json: () => Promise.resolve({ condition: 'Rain', temperature: 'moderate' }),
    })
  ) as jest.Mock;

  render(<Weather />);
  
  fireEvent.change(screen.getByPlaceholderText('Latitude'), { target: { value: '35' } });
  fireEvent.change(screen.getByPlaceholderText('Longitude'), { target: { value: '139' } });
  fireEvent.click(screen.getByText('Check Weather'));

  expect(await screen.findByText('Condition: Rain')).toBeInTheDocument();
  expect(await screen.findByText('Temperature: moderate')).toBeInTheDocument();

  (global.fetch as jest.Mock).mockRestore();
});
