import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VoteitemComponent } from './voteitem.component';

describe('VoteitemComponent', () => {
  let component: VoteitemComponent;
  let fixture: ComponentFixture<VoteitemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VoteitemComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(VoteitemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
