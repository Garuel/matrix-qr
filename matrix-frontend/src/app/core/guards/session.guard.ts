import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth/auth.service';

export const SesionIniciadaGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);
  const router = inject(Router);

  return validate(authService, router);
};

const validate = (authService: AuthService, router: Router) => {
  if (authService.estaLoggeado()) {
    return true;
  }

  return router.createUrlTree(['auth']);
};
