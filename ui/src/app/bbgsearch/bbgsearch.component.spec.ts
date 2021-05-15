import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BbgsearchComponent } from './bbgsearch.component';

describe('BbgsearchComponent', () => {
  let component: BbgsearchComponent;
  let fixture: ComponentFixture<BbgsearchComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BbgsearchComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BbgsearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
