import {Component, OnInit}          from 'angular2/core';
import {BlacklistEntry}             from './blacklist-entry';
import {BlacklistListComponent}     from './blacklist-list.component';
import {BlacklistDetailComponent}   from './blacklist-detail.component';
import {BlacklistService}           from './blacklist.service';

@Component({
    selector: 'blacklist',
    template: `
    <!-- navbar -->
	<div class="row">
      <div class="col-lg-4">
        <blacklist-list [entries]="blacklist_entries"></blacklist-list>
      </div>
      <div class="col-lg-4">
        <blacklist-details [entry]="selectedEntry"></blacklist-details>
      </div>
    </div>
    `,
    directives: [BlacklistListComponent, BlacklistDetailComponent]
})

export class BlacklistComponent implements OnInit {
    title = 'Blacklist';
    blacklist_entries: BlacklistEntry[];

    constructor(private _blacklistService: BlacklistService) { }

    getBlacklist() {
        this._blacklistService.getEntries()
            .then(blacklist_entries => this.blacklist_entries = blacklist_entries);
    }

    ngOnInit() {
        this.getBlacklist();
    }
}
