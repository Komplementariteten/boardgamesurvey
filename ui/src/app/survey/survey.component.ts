import { Component, OnInit } from '@angular/core';
import { ListService } from '../list.service';
import { SurveyItem } from '../surveyitem';

@Component({
  selector: 'app-survey',
  templateUrl: './survey.component.html',
  styleUrls: ['./survey.component.css']
})
export class SurveyComponent implements OnInit {

  showSearch: boolean = false;
  suggestGame: boolean = true;
  constructor(public list:ListService) {
    let sItem = new SurveyItem();
    sItem.Title = "test this is a longer text which should break at some point so this keeps going on and on and on";
    sItem.Description = "someDesc";
    this.list.add(sItem);
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
