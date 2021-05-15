import { TestBed } from '@angular/core/testing';

import { BbgService } from './bbg.service';

describe('BbgService', () => {
  let service: BbgService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BbgService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
