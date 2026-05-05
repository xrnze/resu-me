import React, { createContext } from 'react';

interface ThemeContextType {
  theme: 'dark' | 'light';
  toggle: () => void;
}

export const ThemeContext = createContext<ThemeContextType>({
  theme: 'dark',
  toggle: () => {},
});

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [theme, setTheme] = React.useState<'dark' | 'light'>('light');

  const toggle = () => {
    setTheme(prev => prev === 'dark' ? 'light' : 'dark');
  };

  return (
    <ThemeContext.Provider value={{ theme, toggle }}>
      {children}
    </ThemeContext.Provider>
  );
}
