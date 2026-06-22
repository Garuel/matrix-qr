import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';
import { SesionIniciadaGuard } from './session.guard';

describe('SesionIniciadaGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) =>
    TestBed.runInInjectionContext(() => SesionIniciadaGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
