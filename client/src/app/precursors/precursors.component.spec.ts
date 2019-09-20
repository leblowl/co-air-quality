import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PrecursorsComponent } from './precursors.component';

describe('PrecursorsComponent', () => {
  let component: PrecursorsComponent;
  let fixture: ComponentFixture<PrecursorsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PrecursorsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PrecursorsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
