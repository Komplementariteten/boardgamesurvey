import { Component, OnInit } from '@angular/core';
import { BbgService } from '../bbg.service'

@Component({
  selector: 'app-survey',
  templateUrl: "./survey.component.html",
  styles: [
  ]
})
export class SurveyComponent implements OnInit {

  constructor(private bbgService: BbgService) { }

  ngOnInit(): void {
    this.bbgService.search("abc").subscribe(d => {
      console.log(d);
    });
  }

}
