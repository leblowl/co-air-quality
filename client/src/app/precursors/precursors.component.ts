import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-precursors',
  templateUrl: './precursors.component.html',
  styleUrls: ['./precursors.component.scss']
})
export class PrecursorsComponent implements OnInit {
  units = {
    ppbC: 'ppbC'
  };

  precursors = [
    {
      site: 'PVCO',
      sample_date: '1/8/2018',
      analysis_date: '1/16/2018',
      analyte: 'Propane',
      cas_number: '74-98-6',
      result: '185',
      qualifier: '',
      detection_limit: '0.12',
      units: this.units.ppbC
    },
    {
      site: 'PVCO',
      sample_date: '1/8/2018',
      analysis_date: '1/16/2018',
      analyte: 'Propylene',
      cas_number: '115-07-1',
      result: '1.81',
      qualifier: 'J',
      detection_limit: '0.1',
      units: this.units.ppbC
    }
  ];

  constructor() { }

  ngOnInit() {
  }

}
