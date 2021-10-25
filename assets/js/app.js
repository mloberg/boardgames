import { infiniteHits, searchBox, configure, sortBy, toggleRefinement } from 'instantsearch.js/es/widgets';
import instantsearch from 'instantsearch.js';
import { meiliHost, meiliApiKey } from '@params'; // eslint-disable-line import/no-unresolved
import { instantMeiliSearch } from '@meilisearch/instant-meilisearch';

const indexName = 'boardgames';
const search = instantsearch({
    indexName,
    searchClient: instantMeiliSearch(meiliHost, meiliApiKey),
});

search.addWidgets([
    configure({
        hitsPerPage: 24,
    }),
    searchBox({
        container: '#searchbox',
        showSubmit: false,
        placeholder: 'Search',
        cssClasses: {
            input: 'border border-2 ring-1 pl-2 py-1',
            reset: '-ml-6',
        },
    }),
    sortBy({
        container: '#sort-by',
        items: [
            { label: 'Relevant', value: indexName },
            { label: 'Name', value: `${indexName}:name:asc` },
            { label: 'Rating', value: `${indexName}:rating:desc` },
            { label: 'Complexity', value: `${indexName}:weight:desc` },
        ],
        cssClasses: {
            select: 'border border-2 ring-1 pl-2 py-1',
        },
    }),
    toggleRefinement({
        container: '#expansions',
        attribute: 'type',
        on: 'boardgame',
        cssClasses: {
            checkbox: 'mr-2',
        },
        templates: {
            labelText: 'Hide Expansions',
        },
    }),
    infiniteHits({
        container: '#hits',
        transformItems(items) {
            return items.map((item) => {
                const { minplaytime, maxplaytime, minplayers, maxplayers, rating, weight } = item;
                return {
                    ...item,
                    rating: rating.toFixed(1),
                    weight: weight.toFixed(2),
                    playtime: minplaytime === maxplaytime ? maxplaytime : `${minplaytime}-${maxplaytime}`,
                    players: minplayers === maxplayers ? maxplayers : `${minplayers}-${maxplayers}`,
                };
            });
        },
        cssClasses: {
            list: 'grid md:grid-cols-2 lg:grid-cols-3 gap-4',
            loadMore:
                'mt-6 p-2 pl-5 pr-5 bg-blue-500 hover:bg-blue-700 text-gray-100 text-lg rounded-lg focus:border-4 border-blue-300',
            disabledLoadMore: 'cursor-not-allowed',
        },
        templates: {
            item: `
                <div class="rounded bg-white shadow-md border">
                    <div class="flex">
                        <div class="mx-5 relative">
                                <img class="my-5 rounded-sm" src="{{ thumbnail }}" alt="{{ name }}">
                                <span class="absolute top-3 -left-2 h-10 w-10 bg-blue-700 bg-opacity-75 rounded-full text-white text-lg font-bold flex flex-col justify-center items-center">{{ rating }}</span>
                        </div>
                        <div class="mt-4 mr-4 ml-4 md:ml-0">
                            <p class="text-xl font-bold">{{#helpers.highlight}}{ "attribute": "name" }{{/helpers.highlight}}</p>
                            <p>{{ players }} Players | {{ playtime }} minutes</p>
                            <p><strong>Complexity</strong>: {{ weight }} / 5</p>
                            <p>
                                <a class="text-blue-700 hover:text-blue-900 inline-flex" href="https://boardgamegeek.com/boardgame/{{ id }}" target="_blank" rel="noopener">
                                    More information
                                    <svg class="w-4 h-4 ml-1 mt-1 fill-current" viewBox="0 0 24 24">
                                        <path d="M19 6.41L8.7 16.71a1 1 0 1 1-1.4-1.42L17.58 5H14a1 1 0 0 1 0-2h6a1 1 0 0 1 1 1v6a1 1 0 0 1-2 0V6.41zM17 14a1 1 0 0 1 2 0v5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7c0-1.1.9-2 2-2h5a1 1 0 0 1 0 2H5v12h12v-5z"></path>
                                    </svg>
                                </a>
                            </p>
                        </div>
                    </div>
                </div>
            `,
        },
    }),
    {
        init({ helper }) {
            const players = document.querySelector('#players');
            const filterPlayers = () => {
                const value = players.value;
                helper.removeNumericRefinement('maxplayers');
                helper.removeNumericRefinement('minplayers');

                if (value) {
                    helper.addNumericRefinement('maxplayers', '>=', value);
                    helper.addNumericRefinement('minplayers', '<=', value);
                }

                helper.search();
            };
            players.addEventListener('keyup', filterPlayers);
            players.addEventListener('mouseup', filterPlayers);

            const playtime = document.querySelector('#playtime');
            const filterPlayTime = () => {
                const value = playtime.value;
                helper.removeNumericRefinement('maxplaytime');

                if (value) {
                    helper.addNumericRefinement('maxplaytime', '<=', value);
                }

                helper.search();
            };
            playtime.addEventListener('keyup', filterPlayTime);
            playtime.addEventListener('mouseup', filterPlayTime);
        },
    },
]);

search.start();
