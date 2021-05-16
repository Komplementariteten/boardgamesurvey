import { Component, OnInit } from '@angular/core';
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

  public onSearch(event: Event) {
    if (this.searchText.length > 3) {
      this.isSearching = true
      this.bbgService.search(this.searchText).subscribe(i => {
        this.items = i;
        this.isSearching = false
      });
    }
  }

}
