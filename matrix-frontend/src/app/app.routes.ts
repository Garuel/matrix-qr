import { Routes } from '@angular/router';
import { SesionIniciadaGuard } from './core/guards/session.guard';

export const routes: Routes = [
  {
    path: 'auth/login',
    loadComponent: () =>
      import('./features/auth/login/login.component').then((m) => m.LoginComponent),
  },

  {
    path: '',
    loadComponent: () =>
      import('./shared/components/layout/layout.component').then((m) => m.LayoutComponent),
    canActivate: [SesionIniciadaGuard],
    children: [
      {
        path: 'matrix',
        loadComponent: () =>
          import('./features/matrix/matrix.component').then((m) => m.MatrixComponent),
      },
      { path: '', redirectTo: 'matrix', pathMatch: 'full' },
    ],
  },

  { path: '**', redirectTo: 'auth/login' },
];
