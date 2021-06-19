import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import {BbgService} from '../bbg.service'
import { BoardGame } from '../boardgame';
import { SearchItem } from '../searchitem';

@Component({
  selector: 'app-bbgsearch',
  templateUrl: './bbgsearch.component.html',
  styleUrls: ['./bbgsearch.component.css']
})
export class BbgsearchComponent implements OnInit {

  searchText: string
  items: SearchItem[] | undefined
  isSearching: boolean = false
  searchHasError: boolean = false;
  selectedGame: BoardGame | undefined
  selectionIsShown: boolean = false
  constructor(private bbgService: BbgService) {
    this.searchText = ""
  }

  ngOnInit(): void {

  }

  public onItemSelect(item: SearchItem) {
    console.log(item);
    this.isSearching = true
    this.bbgService.getDetails(item.ObjectId, this.searchText.toLowerCase()).subscribe(d => {
      this.selectedGame = d;
      this.isSearching = false;
      this.selectionIsShown = true;
    });
  }
  private handleError(error: HttpErrorResponse) {
    this.searchHasError = true;
    console.log("search Error");
    return [];
  }
  public onSearch(event: Event) {
    if (this.searchText.length > 3) {
      this.isSearching = true
      this.bbgService.search(this.searchText).subscribe(i => {
        if(i.length > 0){
          this.items = i;
        } else {
          this.searchHasError = true;
        }
        this.isSearching = false
      });
    }
  }

}
