import { Component, OnInit } from '@angular/core';
import { Session } from '../session';
import { ListService } from '../list.service';
import { SessionService } from '../session.service';
import { SurveyItem } from '../surveyitem';

@Component({
  selector: 'app-survey',
  templateUrl: './survey.component.html',
  styleUrls: ['./survey.component.css']
})
export class SurveyComponent implements OnInit {

  showSearch: boolean = false;
  suggestGame: boolean = true;
  constructor(public list:ListService, private sessionService: SessionService) {

    this.sessionService.loadInitial().subscribe(s => {
      this.list.initialize(s);
    })
  }

  ngOnInit(): void {
  }
  showSuggestion() {
    this.showSearch = true;
  }
  hideSuggestion() {
    this.showSearch = false;
  }
}
