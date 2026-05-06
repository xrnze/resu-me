import { Outlet } from '@tanstack/react-router';
import { ThemeProvider } from '@/contexts/ThemeContext';

export function RootLayout() {
  return (
    <ThemeProvider>
      <Outlet />
    </ThemeProvider>
  );
}
